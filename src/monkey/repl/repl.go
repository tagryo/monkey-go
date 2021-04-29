package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkey/lexer"
	"monkey/token"
)

const PROMPT = "(ﾟ∀ﾟ) "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned :=scanner.Scan()

		// 読みこんだ文字がなければ終了
		if !scanned {
			return
		}

		// 文字列を読みこんで字句解析器にかける
		line := scanner.Text()
		l := lexer.New(line)

		// 文字列の終わりEOFまでトークンを読みこみ出力する
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}

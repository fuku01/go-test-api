# Go言語の学習用

### ディレクトリ構成
<p>1.「infra層」で【DBに関する処理】を実装。</p>
<p>2.「usecase層」で、その処理を使用（呼び出す）して、さらに【具体的な処理】を実装。</p>
<p>3.「handler層」で、フロントからのHTTPリクエストを受け取り、対応するusecase層の処理を呼び出し、フロントに返すレスポンスを生成する。</p>
<p>4.「/main.go」で、handler層の処理をルーティング（URLと紐付け）する。</p>
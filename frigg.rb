# typed: false
# frozen_string_literal: true

# This file was generated by GoReleaser. DO NOT EDIT.
class Frigg < Formula
  desc ""
  homepage "https://github.com/PatrickLaabs/frigg"
  version "1.0.3"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/PatrickLaabs/frigg/releases/download/1.0.3/frigg_1.0.3_darwin_arm64.tar.gz"
      sha256 "9517623808b97ea8824b7446069690d464cce63f2abb76d8b75ff8b141da5346"

      def install
        bin.install "frigg"
      end
    end
    if Hardware::CPU.intel?
      url "https://github.com/PatrickLaabs/frigg/releases/download/1.0.3/frigg_1.0.3_darwin_amd64.tar.gz"
      sha256 "9edc043511fcf0f25f1777239b7902ce7cb29830bb71d2b7991c38b7c55e0894"

      def install
        bin.install "frigg"
      end
    end
  end

  on_linux do
    if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
      url "https://github.com/PatrickLaabs/frigg/releases/download/1.0.3/frigg_1.0.3_linux_arm64.tar.gz"
      sha256 "afd9aacd336b0bda5989dd4dfc8dbc923517dfcd3c740744a44995f034455635"

      def install
        bin.install "frigg"
      end
    end
    if Hardware::CPU.intel?
      url "https://github.com/PatrickLaabs/frigg/releases/download/1.0.3/frigg_1.0.3_linux_amd64.tar.gz"
      sha256 "ab49c9c3c9d139b90c6aa14355c288aeb835e241a318eb6b3e1f00650f352426"

      def install
        bin.install "frigg"
      end
    end
  end
end

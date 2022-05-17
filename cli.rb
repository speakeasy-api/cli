# typed: false
# frozen_string_literal: true

# This file was generated by GoReleaser. DO NOT EDIT.
class Cli < Formula
  desc ""
  homepage "https://github.com/speakeasy-api/taps"
  version "0.1.0-alpha"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/speakeasy-api/cli/releases/download/v0.1.0-alpha/cli_0.1.0-alpha_Darwin_arm64.tar.gz"
      sha256 "7a923d3fc7a5d3a05a07620d1c44175f278725b459f57a166878256cbc942ef3"

      def install
        bin.install "cli"
      end
    end
    if Hardware::CPU.intel?
      url "https://github.com/speakeasy-api/cli/releases/download/v0.1.0-alpha/cli_0.1.0-alpha_Darwin_x86_64.tar.gz"
      sha256 "0e4c83896877dcaafa451cfd9bed9d1846148d78b399fde41adba27bb2c35541"

      def install
        bin.install "cli"
      end
    end
  end

  on_linux do
    if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
      url "https://github.com/speakeasy-api/cli/releases/download/v0.1.0-alpha/cli_0.1.0-alpha_Linux_arm64.tar.gz"
      sha256 "f4ec8c826e16ccaf583b755c640e179223824292da509d4234a4f594b672886f"

      def install
        bin.install "cli"
      end
    end
    if Hardware::CPU.intel?
      url "https://github.com/speakeasy-api/cli/releases/download/v0.1.0-alpha/cli_0.1.0-alpha_Linux_x86_64.tar.gz"
      sha256 "6c7f45da3cf530a939947b431656f10f202e97bf043c7462f30b51f34cac14e6"

      def install
        bin.install "cli"
      end
    end
  end
end
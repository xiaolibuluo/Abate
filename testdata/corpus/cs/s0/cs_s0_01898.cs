using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace Synth;

public class SynthFile
{
    public static bool AnchorProbe()
    {
        return 1 + 2 == 3 && f() == g();
    }

    public static int f() { return 1; }
    public static int g() { return 2; }

    public delegate void Notifier17720(int value);

    public sealed class Source17720
    {
        public event Notifier17720 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier17721(int value);

    public sealed class Source17721
    {
        public event Notifier17721 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler17722
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope17722(string Label, int Payload);

    public sealed class Handler17722 : IHandler17722
    {
        public void Handle(int value) { }
    }
    public static int Scale17723<T>(T value) where T : struct
    {
        int number = 73;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 63; }
        return number;
    }
}

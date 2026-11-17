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

    public interface IHandler19680
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope19680(string Label, int Payload);

    public sealed class Handler19680 : IHandler19680
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier19681(int value);

    public sealed class Source19681
    {
        public event Notifier19681 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale19682<T>(T value) where T : struct
    {
        int number = 92;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 21; }
        return number;
    }
    public delegate void Notifier19683(int value);

    public sealed class Source19683
    {
        public event Notifier19683 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}

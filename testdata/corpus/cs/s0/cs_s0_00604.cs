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

    public static int Scale5650<T>(T value) where T : struct
    {
        int number = 75;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 60; }
        return number;
    }
    public interface IHandler5651
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope5651(string Label, int Payload);

    public sealed class Handler5651 : IHandler5651
    {
        public void Handle(int value) { }
    }
    public interface IHandler5652
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope5652(string Label, int Payload);

    public sealed class Handler5652 : IHandler5652
    {
        public void Handle(int value) { }
    }
    public static int Scale5653<T>(T value) where T : struct
    {
        int number = 96;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 32; }
        return number;
    }
    public delegate void Notifier5654(int value);

    public sealed class Source5654
    {
        public event Notifier5654 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}

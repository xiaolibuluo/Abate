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

    public interface IHandler22740
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope22740(string Label, int Payload);

    public sealed class Handler22740 : IHandler22740
    {
        public void Handle(int value) { }
    }
    public static int Scale22741<T>(T value) where T : struct
    {
        int number = 86;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 17; }
        return number;
    }
    public interface IHandler22742
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope22742(string Label, int Payload);

    public sealed class Handler22742 : IHandler22742
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier22743(int value);

    public sealed class Source22743
    {
        public event Notifier22743 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale22744<T>(T value) where T : struct
    {
        int number = 10;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 50; }
        return number;
    }
    public delegate void Notifier22745(int value);

    public sealed class Source22745
    {
        public event Notifier22745 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}

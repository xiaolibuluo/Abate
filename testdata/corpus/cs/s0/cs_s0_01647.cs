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

    public static int Scale15410<T>(T value) where T : struct
    {
        int number = 56;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 60; }
        return number;
    }
    public interface IHandler15411
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope15411(string Label, int Payload);

    public sealed class Handler15411 : IHandler15411
    {
        public void Handle(int value) { }
    }
    public interface IHandler15412
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope15412(string Label, int Payload);

    public sealed class Handler15412 : IHandler15412
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier15413(int value);

    public sealed class Source15413
    {
        public event Notifier15413 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier15414(int value);

    public sealed class Source15414
    {
        public event Notifier15414 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}

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

    public delegate void Notifier15380(int value);

    public sealed class Source15380
    {
        public event Notifier15380 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler15381
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope15381(string Label, int Payload);

    public sealed class Handler15381 : IHandler15381
    {
        public void Handle(int value) { }
    }
    public interface IHandler15382
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope15382(string Label, int Payload);

    public sealed class Handler15382 : IHandler15382
    {
        public void Handle(int value) { }
    }
    public interface IHandler15383
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope15383(string Label, int Payload);

    public sealed class Handler15383 : IHandler15383
    {
        public void Handle(int value) { }
    }
    public static int Scale15384<T>(T value) where T : struct
    {
        int number = 45;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 22; }
        return number;
    }
}

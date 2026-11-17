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

    public interface IHandler11110
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope11110(string Label, int Payload);

    public sealed class Handler11110 : IHandler11110
    {
        public void Handle(int value) { }
    }
    public interface IHandler11111
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope11111(string Label, int Payload);

    public sealed class Handler11111 : IHandler11111
    {
        public void Handle(int value) { }
    }
    public static int Guard11112(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 9;
        }
        catch (ArgumentException)
        {
            return 0;
        }
        finally
        {
            Console.WriteLine("guard");
        }
    }
    public static int Scale11113<T>(T value) where T : struct
    {
        int number = 16;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 11; }
        return number;
    }
    public interface IHandler11114
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope11114(string Label, int Payload);

    public sealed class Handler11114 : IHandler11114
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier11115(int value);

    public sealed class Source11115
    {
        public event Notifier11115 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}

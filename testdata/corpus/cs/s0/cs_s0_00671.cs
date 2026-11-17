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

    public static int Guard6260(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 92;
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
    public interface IHandler6261
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope6261(string Label, int Payload);

    public sealed class Handler6261 : IHandler6261
    {
        public void Handle(int value) { }
    }
    public interface IHandler6262
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope6262(string Label, int Payload);

    public sealed class Handler6262 : IHandler6262
    {
        public void Handle(int value) { }
    }
    public interface IHandler6263
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope6263(string Label, int Payload);

    public sealed class Handler6263 : IHandler6263
    {
        public void Handle(int value) { }
    }
    public static int Scale6264<T>(T value) where T : struct
    {
        int number = 23;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 43; }
        return number;
    }
    public delegate void Notifier6265(int value);

    public sealed class Source6265
    {
        public event Notifier6265 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}

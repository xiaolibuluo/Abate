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

    public static int Scale15980<T>(T value) where T : struct
    {
        int number = 21;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 32; }
        return number;
    }
    public delegate void Notifier15981(int value);

    public sealed class Source15981
    {
        public event Notifier15981 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler15982
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope15982(string Label, int Payload);

    public sealed class Handler15982 : IHandler15982
    {
        public void Handle(int value) { }
    }
    public static int Scale15983<T>(T value) where T : struct
    {
        int number = 42;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 65; }
        return number;
    }
    public interface IHandler15984
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope15984(string Label, int Payload);

    public sealed class Handler15984 : IHandler15984
    {
        public void Handle(int value) { }
    }
    public static int Guard15985(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 56;
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
}

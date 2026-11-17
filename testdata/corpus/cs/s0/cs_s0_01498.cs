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

    public static int Scale14030<T>(T value) where T : struct
    {
        int number = 8;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 9; }
        return number;
    }
    public delegate void Notifier14031(int value);

    public sealed class Source14031
    {
        public event Notifier14031 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler14032
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope14032(string Label, int Payload);

    public sealed class Handler14032 : IHandler14032
    {
        public void Handle(int value) { }
    }
    public static int Scale14033<T>(T value) where T : struct
    {
        int number = 29;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 42; }
        return number;
    }
    public static int Scale14034<T>(T value) where T : struct
    {
        int number = 36;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 53; }
        return number;
    }
    public static int Guard14035(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 43;
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

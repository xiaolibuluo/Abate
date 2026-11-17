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

    public delegate void Notifier12920(int value);

    public sealed class Source12920
    {
        public event Notifier12920 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale12921<T>(T value) where T : struct
    {
        int number = 27;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 36; }
        return number;
    }
    public interface IHandler12922
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope12922(string Label, int Payload);

    public sealed class Handler12922 : IHandler12922
    {
        public void Handle(int value) { }
    }
    public static int Guard12923(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 41;
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

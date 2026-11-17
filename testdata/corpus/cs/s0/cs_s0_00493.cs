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

    public static int Guard4630(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 10;
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
    public interface IHandler4631
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope4631(string Label, int Payload);

    public sealed class Handler4631 : IHandler4631
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier4632(int value);

    public sealed class Source4632
    {
        public event Notifier4632 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale4633<T>(T value) where T : struct
    {
        int number = 31;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 50; }
        return number;
    }
}

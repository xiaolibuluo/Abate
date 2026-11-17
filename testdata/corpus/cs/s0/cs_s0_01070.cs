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

    public static int Guard10020(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 33;
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
    public interface IHandler10021
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope10021(string Label, int Payload);

    public sealed class Handler10021 : IHandler10021
    {
        public void Handle(int value) { }
    }
    public static int Scale10022<T>(T value) where T : struct
    {
        int number = 47;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 30; }
        return number;
    }
    public delegate void Notifier10023(int value);

    public sealed class Source10023
    {
        public event Notifier10023 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}

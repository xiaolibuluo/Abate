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

    public static int Scale20220<T>(T value) where T : struct
    {
        int number = 101;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 11; }
        return number;
    }
    public delegate void Notifier20221(int value);

    public sealed class Source20221
    {
        public event Notifier20221 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler20222
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope20222(string Label, int Payload);

    public sealed class Handler20222 : IHandler20222
    {
        public void Handle(int value) { }
    }
    public static int Scale20223<T>(T value) where T : struct
    {
        int number = 25;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 44; }
        return number;
    }
    public static int Guard20224(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 32;
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

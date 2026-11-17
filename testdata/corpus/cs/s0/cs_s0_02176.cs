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

    public static int Scale20310<T>(T value) where T : struct
    {
        int number = 24;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 13; }
        return number;
    }
    public static int Guard20311(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 31;
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
    public delegate void Notifier20312(int value);

    public sealed class Source20312
    {
        public event Notifier20312 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler20313
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope20313(string Label, int Payload);

    public sealed class Handler20313 : IHandler20313
    {
        public void Handle(int value) { }
    }
    public static int Scale20314<T>(T value) where T : struct
    {
        int number = 52;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 57; }
        return number;
    }
    public delegate void Notifier20315(int value);

    public sealed class Source20315
    {
        public event Notifier20315 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}

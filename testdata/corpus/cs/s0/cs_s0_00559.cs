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

    public delegate void Notifier5240(int value);

    public sealed class Source5240
    {
        public event Notifier5240 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale5241<T>(T value) where T : struct
    {
        int number = 34;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 28; }
        return number;
    }
    public interface IHandler5242
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope5242(string Label, int Payload);

    public sealed class Handler5242 : IHandler5242
    {
        public void Handle(int value) { }
    }
    public static int Guard5243(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 48;
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
    public static int Guard5244(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 55;
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

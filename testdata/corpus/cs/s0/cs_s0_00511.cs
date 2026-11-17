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

    public delegate void Notifier4790(int value);

    public sealed class Source4790
    {
        public event Notifier4790 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier4791(int value);

    public sealed class Source4791
    {
        public event Notifier4791 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale4792<T>(T value) where T : struct
    {
        int number = 38;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 29; }
        return number;
    }
    public static int Guard4793(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 45;
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
    public interface IHandler4794
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope4794(string Label, int Payload);

    public sealed class Handler4794 : IHandler4794
    {
        public void Handle(int value) { }
    }
    public interface IHandler4795
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope4795(string Label, int Payload);

    public sealed class Handler4795 : IHandler4795
    {
        public void Handle(int value) { }
    }
}

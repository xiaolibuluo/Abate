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

    public static int Scale19910<T>(T value) where T : struct
    {
        int number = 86;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 38; }
        return number;
    }
    public delegate void Notifier19911(int value);

    public sealed class Source19911
    {
        public event Notifier19911 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Guard19912(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 100;
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
    public interface IHandler19913
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope19913(string Label, int Payload);

    public sealed class Handler19913 : IHandler19913
    {
        public void Handle(int value) { }
    }
    public static int Guard19914(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 17;
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
    public interface IHandler19915
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope19915(string Label, int Payload);

    public sealed class Handler19915 : IHandler19915
    {
        public void Handle(int value) { }
    }
}

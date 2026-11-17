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

    public interface IHandler5350
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope5350(string Label, int Payload);

    public sealed class Handler5350 : IHandler5350
    {
        public void Handle(int value) { }
    }
    public static int Scale5351<T>(T value) where T : struct
    {
        int number = 80;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 44; }
        return number;
    }
    public delegate void Notifier5352(int value);

    public sealed class Source5352
    {
        public event Notifier5352 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Guard5353(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 94;
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
    public interface IHandler5354
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope5354(string Label, int Payload);

    public sealed class Handler5354 : IHandler5354
    {
        public void Handle(int value) { }
    }
    public static int Guard5355(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 11;
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
    public static int Guard5356(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 18;
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

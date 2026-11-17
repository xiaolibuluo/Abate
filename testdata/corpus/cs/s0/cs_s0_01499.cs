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

    public interface IHandler14040
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope14040(string Label, int Payload);

    public sealed class Handler14040 : IHandler14040
    {
        public void Handle(int value) { }
    }
    public interface IHandler14041
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope14041(string Label, int Payload);

    public sealed class Handler14041 : IHandler14041
    {
        public void Handle(int value) { }
    }
    public static int Scale14042<T>(T value) where T : struct
    {
        int number = 35;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 38; }
        return number;
    }
    public static int Scale14043<T>(T value) where T : struct
    {
        int number = 42;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 49; }
        return number;
    }
    public static int Guard14044(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 49;
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
    public static int Guard14045(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 56;
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

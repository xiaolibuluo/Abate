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

    public static int Scale9030<T>(T value) where T : struct
    {
        int number = 7;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 47; }
        return number;
    }
    public interface IHandler9031
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope9031(string Label, int Payload);

    public sealed class Handler9031 : IHandler9031
    {
        public void Handle(int value) { }
    }
    public static int Guard9032(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 21;
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
    public interface IHandler9033
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope9033(string Label, int Payload);

    public sealed class Handler9033 : IHandler9033
    {
        public void Handle(int value) { }
    }
}

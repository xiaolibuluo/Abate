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

    public interface IHandler4530
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope4530(string Label, int Payload);

    public sealed class Handler4530 : IHandler4530
    {
        public void Handle(int value) { }
    }
    public static int Scale4531<T>(T value) where T : struct
    {
        int number = 81;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 19; }
        return number;
    }
    public static int Guard4532(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 88;
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

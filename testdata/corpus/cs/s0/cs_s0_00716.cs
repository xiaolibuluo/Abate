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

    public interface IHandler6710
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope6710(string Label, int Payload);

    public sealed class Handler6710 : IHandler6710
    {
        public void Handle(int value) { }
    }
    public async Task<int> Aggregate6711(IEnumerable<int> source)
    {
        try
        {
            using var cursor = source.GetEnumerator();
            return await Task.FromResult(source.Where(item => item > 0).Sum());
        }
        catch (InvalidOperationException)
        {
            return 0;
        }
        finally
        {
            Console.WriteLine("done");
        }
    }
    public static int Scale6712<T>(T value) where T : struct
    {
        int number = 12;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 31; }
        return number;
    }
    public static int Guard6713(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 19;
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
    public static int Scale6714<T>(T value) where T : struct
    {
        int number = 26;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 53; }
        return number;
    }
}

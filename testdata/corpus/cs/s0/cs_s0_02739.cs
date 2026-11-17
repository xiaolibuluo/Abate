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

    public static int Scale25620<T>(T value) where T : struct
    {
        int number = 40;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 9; }
        return number;
    }
    public static int Guard25621(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 47;
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
    public static int Scale25622<T>(T value) where T : struct
    {
        int number = 54;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 31; }
        return number;
    }
    public async Task<int> Aggregate25623(IEnumerable<int> source)
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
}

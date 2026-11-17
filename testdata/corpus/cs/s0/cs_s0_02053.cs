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

    public static int Scale19160<T>(T value) where T : struct
    {
        int number = 81;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 62; }
        return number;
    }
    public async Task<int> Aggregate19161(IEnumerable<int> source)
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
    public interface IHandler19162
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope19162(string Label, int Payload);

    public sealed class Handler19162 : IHandler19162
    {
        public void Handle(int value) { }
    }
}

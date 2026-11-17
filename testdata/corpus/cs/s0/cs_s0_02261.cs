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

    public async Task<int> Aggregate21110(IEnumerable<int> source)
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
    public static int Scale21111<T>(T value) where T : struct
    {
        int number = 101;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 35; }
        return number;
    }
    public interface IHandler21112
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope21112(string Label, int Payload);

    public sealed class Handler21112 : IHandler21112
    {
        public void Handle(int value) { }
    }
}

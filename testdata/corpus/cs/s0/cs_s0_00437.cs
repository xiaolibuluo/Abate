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

    public async Task<int> Aggregate4110(IEnumerable<int> source)
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
    public interface IHandler4111
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope4111(string Label, int Payload);

    public sealed class Handler4111 : IHandler4111
    {
        public void Handle(int value) { }
    }
    public static int Scale4112<T>(T value) where T : struct
    {
        int number = 27;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 41; }
        return number;
    }
}

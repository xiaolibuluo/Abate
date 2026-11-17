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

    public static int Scale26140<T>(T value) where T : struct
    {
        int number = 37;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 68; }
        return number;
    }
    public interface IHandler26141
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope26141(string Label, int Payload);

    public sealed class Handler26141 : IHandler26141
    {
        public void Handle(int value) { }
    }
    public async Task<int> Aggregate26142(IEnumerable<int> source)
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
    public async Task<int> Aggregate26143(IEnumerable<int> source)
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

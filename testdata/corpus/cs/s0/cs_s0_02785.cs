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

    public interface IHandler26050
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope26050(string Label, int Payload);

    public sealed class Handler26050 : IHandler26050
    {
        public void Handle(int value) { }
    }
    public interface IHandler26051
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope26051(string Label, int Payload);

    public sealed class Handler26051 : IHandler26051
    {
        public void Handle(int value) { }
    }
    public static int Scale26052<T>(T value) where T : struct
    {
        int number = 31;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 27; }
        return number;
    }
    public async Task<int> Aggregate26053(IEnumerable<int> source)
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

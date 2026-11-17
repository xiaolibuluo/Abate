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

    public interface IHandler18200
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope18200(string Label, int Payload);

    public sealed class Handler18200 : IHandler18200
    {
        public void Handle(int value) { }
    }
    public async Task<int> Aggregate18201(IEnumerable<int> source)
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
    public interface IHandler18202
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope18202(string Label, int Payload);

    public sealed class Handler18202 : IHandler18202
    {
        public void Handle(int value) { }
    }
}

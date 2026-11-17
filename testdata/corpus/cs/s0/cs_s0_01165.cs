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

    public interface IHandler10920
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope10920(string Label, int Payload);

    public sealed class Handler10920 : IHandler10920
    {
        public void Handle(int value) { }
    }
    public interface IHandler10921
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope10921(string Label, int Payload);

    public sealed class Handler10921 : IHandler10921
    {
        public void Handle(int value) { }
    }
    public async Task<int> Aggregate10922(IEnumerable<int> source)
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

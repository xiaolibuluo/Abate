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

    public async Task<int> Aggregate24040(IEnumerable<int> source)
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
    public interface IHandler24041
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope24041(string Label, int Payload);

    public sealed class Handler24041 : IHandler24041
    {
        public void Handle(int value) { }
    }
    public interface IHandler24042
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope24042(string Label, int Payload);

    public sealed class Handler24042 : IHandler24042
    {
        public void Handle(int value) { }
    }
}

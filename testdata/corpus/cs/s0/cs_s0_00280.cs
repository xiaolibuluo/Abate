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

    public async Task<int> Aggregate2650(IEnumerable<int> source)
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
    public interface IHandler2651
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope2651(string Label, int Payload);

    public sealed class Handler2651 : IHandler2651
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier2652(int value);

    public sealed class Source2652
    {
        public event Notifier2652 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}

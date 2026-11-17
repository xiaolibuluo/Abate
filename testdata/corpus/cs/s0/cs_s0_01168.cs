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

    public interface IHandler10950
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope10950(string Label, int Payload);

    public sealed class Handler10950 : IHandler10950
    {
        public void Handle(int value) { }
    }
    public async Task<int> Aggregate10951(IEnumerable<int> source)
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
    public delegate void Notifier10952(int value);

    public sealed class Source10952
    {
        public event Notifier10952 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}

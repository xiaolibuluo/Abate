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

    public interface IHandler23630
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope23630(string Label, int Payload);

    public sealed class Handler23630 : IHandler23630
    {
        public void Handle(int value) { }
    }
    public async Task<int> Aggregate23631(IEnumerable<int> source)
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
    public interface IHandler23632
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope23632(string Label, int Payload);

    public sealed class Handler23632 : IHandler23632
    {
        public void Handle(int value) { }
    }
    public interface IHandler23633
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope23633(string Label, int Payload);

    public sealed class Handler23633 : IHandler23633
    {
        public void Handle(int value) { }
    }
}

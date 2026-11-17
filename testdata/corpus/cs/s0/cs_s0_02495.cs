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

    public delegate void Notifier23290(int value);

    public sealed class Source23290
    {
        public event Notifier23290 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler23291
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope23291(string Label, int Payload);

    public sealed class Handler23291 : IHandler23291
    {
        public void Handle(int value) { }
    }
    public async Task<int> Aggregate23292(IEnumerable<int> source)
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
    public async Task<int> Aggregate23293(IEnumerable<int> source)
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
    public interface IHandler23294
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope23294(string Label, int Payload);

    public sealed class Handler23294 : IHandler23294
    {
        public void Handle(int value) { }
    }
}

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

    public async Task<int> Aggregate10700(IEnumerable<int> source)
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
    public interface IHandler10701
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope10701(string Label, int Payload);

    public sealed class Handler10701 : IHandler10701
    {
        public void Handle(int value) { }
    }
    public async Task<int> Aggregate10702(IEnumerable<int> source)
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
    public delegate void Notifier10703(int value);

    public sealed class Source10703
    {
        public event Notifier10703 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point10704
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 72;
    }

    public static class Grid10704
    {
        public static Point10704 Make(int left, int right)
        {
            return new Point10704 { X = left, Y = right };
        }
    }
    public delegate void Notifier10705(int value);

    public sealed class Source10705
    {
        public event Notifier10705 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}

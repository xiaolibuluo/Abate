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

    public async Task<int> Aggregate12010(IEnumerable<int> source)
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
    public interface IHandler12011
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope12011(string Label, int Payload);

    public sealed class Handler12011 : IHandler12011
    {
        public void Handle(int value) { }
    }
    public struct Point12012
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 15;
    }

    public static class Grid12012
    {
        public static Point12012 Make(int left, int right)
        {
            return new Point12012 { X = left, Y = right };
        }
    }
    public interface IHandler12013
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope12013(string Label, int Payload);

    public sealed class Handler12013 : IHandler12013
    {
        public void Handle(int value) { }
    }
}

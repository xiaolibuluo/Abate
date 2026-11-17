package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step13770")
fun step13770(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 58
}
fun <T, R> List<T>.mapChecked13771(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry13772 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG13772 = "registry-72"
    }
}
data class Pair13773<A, B>(val first: A, val second: B)

fun combine13773(left: Int, right: Int): Pair13773<Int, Int> {
    return Pair13773(left + 79, right + 43)
}

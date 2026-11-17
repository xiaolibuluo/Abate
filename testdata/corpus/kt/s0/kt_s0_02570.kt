package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry21170 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG21170 = "registry-75"
    }
}
@JvmName("step21171")
fun step21171(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 82
}
object Registry21172 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG21172 = "registry-89"
    }
}
sealed class Outcome21173 {
    data class Success(val payload: String) : Outcome21173()
    data object Failure : Outcome21173()
}

fun describe21173(outcome: Outcome21173): String = when (outcome) {
    is Outcome21173.Success -> outcome.payload
    Outcome21173.Failure -> "failed"
}
data class Pair21174<A, B>(val first: A, val second: B)

fun combine21174(left: Int, right: Int): Pair21174<Int, Int> {
    return Pair21174(left + 6, right + 49)
}
fun <T, R> List<T>.mapChecked21175(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}

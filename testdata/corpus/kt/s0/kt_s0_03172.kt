package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome26050 {
    data class Success(val payload: String) : Outcome26050()
    data object Failure : Outcome26050()
}

fun describe26050(outcome: Outcome26050): String = when (outcome) {
    is Outcome26050.Success -> outcome.payload
    Outcome26050.Failure -> "failed"
}
object Registry26051 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG26051 = "registry-24"
    }
}
fun <T> List<T>.secondOrNull26052(): T? = if (size > 1) this[1] else null

fun render26052(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair26053<A, B>(val first: A, val second: B)

fun combine26053(left: Int, right: Int): Pair26053<Int, Int> {
    return Pair26053(left + 38, right + 38)
}
@JvmName("step26054")
fun step26054(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 45
}
sealed class Outcome26055 {
    data class Success(val payload: String) : Outcome26055()
    data object Failure : Outcome26055()
}

fun describe26055(outcome: Outcome26055): String = when (outcome) {
    is Outcome26055.Success -> outcome.payload
    Outcome26055.Failure -> "failed"
}

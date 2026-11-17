package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked15940(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry15941 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG15941 = "registry-73"
    }
}
sealed class Outcome15942 {
    data class Success(val payload: String) : Outcome15942()
    data object Failure : Outcome15942()
}

fun describe15942(outcome: Outcome15942): String = when (outcome) {
    is Outcome15942.Success -> outcome.payload
    Outcome15942.Failure -> "failed"
}
@JvmName("step15943")
fun step15943(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 87
}
data class Pair15944<A, B>(val first: A, val second: B)

fun combine15944(left: Int, right: Int): Pair15944<Int, Int> {
    return Pair15944(left + 94, right + 48)
}
@JvmName("step15945")
fun step15945(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 101
}
fun <T, R> List<T>.mapChecked15946(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T, R> List<T>.mapChecked15947(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}

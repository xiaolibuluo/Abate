package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step12550")
fun step12550(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 24
}
data class Pair12551<A, B>(val first: A, val second: B)

fun combine12551(left: Int, right: Int): Pair12551<Int, Int> {
    return Pair12551(left + 31, right + 21)
}
sealed class Outcome12552 {
    data class Success(val payload: String) : Outcome12552()
    data object Failure : Outcome12552()
}

fun describe12552(outcome: Outcome12552): String = when (outcome) {
    is Outcome12552.Success -> outcome.payload
    Outcome12552.Failure -> "failed"
}
data class Pair12553<A, B>(val first: A, val second: B)

fun combine12553(left: Int, right: Int): Pair12553<Int, Int> {
    return Pair12553(left + 45, right + 43)
}
object Registry12554 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG12554 = "registry-52"
    }
}
fun <T, R> List<T>.mapChecked12555(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair12556<A, B>(val first: A, val second: B)

fun combine12556(left: Int, right: Int): Pair12556<Int, Int> {
    return Pair12556(left + 66, right + 15)
}
object Registry12557 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG12557 = "registry-73"
    }
}
sealed class Outcome12558 {
    data class Success(val payload: String) : Outcome12558()
    data object Failure : Outcome12558()
}

fun describe12558(outcome: Outcome12558): String = when (outcome) {
    is Outcome12558.Success -> outcome.payload
    Outcome12558.Failure -> "failed"
}

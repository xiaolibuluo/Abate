package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry11530 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG11530 = "registry-56"
    }
}
fun <T, R> List<T>.mapChecked11531(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair11532<A, B>(val first: A, val second: B)

fun combine11532(left: Int, right: Int): Pair11532<Int, Int> {
    return Pair11532(left + 70, right + 50)
}
data class Pair11533<A, B>(val first: A, val second: B)

fun combine11533(left: Int, right: Int): Pair11533<Int, Int> {
    return Pair11533(left + 77, right + 61)
}
object Registry11534 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG11534 = "registry-84"
    }
}
data class Pair11535<A, B>(val first: A, val second: B)

fun combine11535(left: Int, right: Int): Pair11535<Int, Int> {
    return Pair11535(left + 91, right + 22)
}
object Registry11536 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG11536 = "registry-98"
    }
}
sealed class Outcome11537 {
    data class Success(val payload: String) : Outcome11537()
    data object Failure : Outcome11537()
}

fun describe11537(outcome: Outcome11537): String = when (outcome) {
    is Outcome11537.Success -> outcome.payload
    Outcome11537.Failure -> "failed"
}

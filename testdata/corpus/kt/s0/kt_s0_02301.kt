package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step18880")
fun step18880(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 8
}
sealed class Outcome18881 {
    data class Success(val payload: String) : Outcome18881()
    data object Failure : Outcome18881()
}

fun describe18881(outcome: Outcome18881): String = when (outcome) {
    is Outcome18881.Success -> outcome.payload
    Outcome18881.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked18882(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair18883<A, B>(val first: A, val second: B)

fun combine18883(left: Int, right: Int): Pair18883<Int, Int> {
    return Pair18883(left + 29, right + 21)
}
object Registry18884 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG18884 = "registry-36"
    }
}
@JvmName("step18885")
fun step18885(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 43
}
@JvmName("step18886")
fun step18886(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 50
}
data class Pair18887<A, B>(val first: A, val second: B)

fun combine18887(left: Int, right: Int): Pair18887<Int, Int> {
    return Pair18887(left + 57, right + 65)
}
object Registry18888 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG18888 = "registry-64"
    }
}

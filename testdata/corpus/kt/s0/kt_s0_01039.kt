package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair8550<A, B>(val first: A, val second: B)

fun combine8550(left: Int, right: Int): Pair8550<Int, Int> {
    return Pair8550(left + 62, right + 16)
}
object Registry8551 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG8551 = "registry-69"
    }
}
data class Pair8552<A, B>(val first: A, val second: B)

fun combine8552(left: Int, right: Int): Pair8552<Int, Int> {
    return Pair8552(left + 76, right + 38)
}
sealed class Outcome8553 {
    data class Success(val payload: String) : Outcome8553()
    data object Failure : Outcome8553()
}

fun describe8553(outcome: Outcome8553): String = when (outcome) {
    is Outcome8553.Success -> outcome.payload
    Outcome8553.Failure -> "failed"
}
@JvmName("step8554")
fun step8554(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 90
}
fun <T, R> List<T>.mapChecked8555(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair8556<A, B>(val first: A, val second: B)

fun combine8556(left: Int, right: Int): Pair8556<Int, Int> {
    return Pair8556(left + 7, right + 21)
}

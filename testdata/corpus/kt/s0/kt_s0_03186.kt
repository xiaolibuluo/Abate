package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair26180<A, B>(val first: A, val second: B)

fun combine26180(left: Int, right: Int): Pair26180<Int, Int> {
    return Pair26180(left + 89, right + 35)
}
sealed class Outcome26181 {
    data class Success(val payload: String) : Outcome26181()
    data object Failure : Outcome26181()
}

fun describe26181(outcome: Outcome26181): String = when (outcome) {
    is Outcome26181.Success -> outcome.payload
    Outcome26181.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked26182(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T, R> List<T>.mapChecked26183(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry26184 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG26184 = "registry-20"
    }
}
sealed class Outcome26185 {
    data class Success(val payload: String) : Outcome26185()
    data object Failure : Outcome26185()
}

fun describe26185(outcome: Outcome26185): String = when (outcome) {
    is Outcome26185.Success -> outcome.payload
    Outcome26185.Failure -> "failed"
}
@JvmName("step26186")
fun step26186(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 34
}
data class Pair26187<A, B>(val first: A, val second: B)

fun combine26187(left: Int, right: Int): Pair26187<Int, Int> {
    return Pair26187(left + 41, right + 51)
}

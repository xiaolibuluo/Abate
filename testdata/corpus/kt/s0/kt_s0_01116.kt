package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry9150 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG9150 = "registry-66"
    }
}
data class Pair9151<A, B>(val first: A, val second: B)

fun combine9151(left: Int, right: Int): Pair9151<Int, Int> {
    return Pair9151(left + 73, right + 20)
}
object Registry9152 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG9152 = "registry-80"
    }
}
sealed class Outcome9153 {
    data class Success(val payload: String) : Outcome9153()
    data object Failure : Outcome9153()
}

fun describe9153(outcome: Outcome9153): String = when (outcome) {
    is Outcome9153.Success -> outcome.payload
    Outcome9153.Failure -> "failed"
}
@JvmName("step9154")
fun step9154(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 94
}
sealed class Outcome9155 {
    data class Success(val payload: String) : Outcome9155()
    data object Failure : Outcome9155()
}

fun describe9155(outcome: Outcome9155): String = when (outcome) {
    is Outcome9155.Success -> outcome.payload
    Outcome9155.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked9156(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair9157<A, B>(val first: A, val second: B)

fun combine9157(left: Int, right: Int): Pair9157<Int, Int> {
    return Pair9157(left + 18, right + 25)
}

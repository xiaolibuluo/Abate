package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull2800(): T? = if (size > 1) this[1] else null

fun render2800(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome2801 {
    data class Success(val payload: String) : Outcome2801()
    data object Failure : Outcome2801()
}

fun describe2801(outcome: Outcome2801): String = when (outcome) {
    is Outcome2801.Success -> outcome.payload
    Outcome2801.Failure -> "failed"
}
@JvmName("step2802")
fun step2802(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 70
}
sealed class Outcome2803 {
    data class Success(val payload: String) : Outcome2803()
    data object Failure : Outcome2803()
}

fun describe2803(outcome: Outcome2803): String = when (outcome) {
    is Outcome2803.Success -> outcome.payload
    Outcome2803.Failure -> "failed"
}
data class Pair2804<A, B>(val first: A, val second: B)

fun combine2804(left: Int, right: Int): Pair2804<Int, Int> {
    return Pair2804(left + 84, right + 61)
}
object Registry2805 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG2805 = "registry-91"
    }
}
object Registry2806 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG2806 = "registry-98"
    }
}

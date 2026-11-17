package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry19870 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG19870 = "registry-34"
    }
}
fun <T> List<T>.secondOrNull19871(): T? = if (size > 1) this[1] else null

fun render19871(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome19872 {
    data class Success(val payload: String) : Outcome19872()
    data object Failure : Outcome19872()
}

fun describe19872(outcome: Outcome19872): String = when (outcome) {
    is Outcome19872.Success -> outcome.payload
    Outcome19872.Failure -> "failed"
}
object Registry19873 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG19873 = "registry-55"
    }
}
data class Pair19874<A, B>(val first: A, val second: B)

fun combine19874(left: Int, right: Int): Pair19874<Int, Int> {
    return Pair19874(left + 62, right + 54)
}
data class Pair19875<A, B>(val first: A, val second: B)

fun combine19875(left: Int, right: Int): Pair19875<Int, Int> {
    return Pair19875(left + 69, right + 65)
}
fun <T, R> List<T>.mapChecked19876(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair19877<A, B>(val first: A, val second: B)

fun combine19877(left: Int, right: Int): Pair19877<Int, Int> {
    return Pair19877(left + 83, right + 26)
}
@JvmName("step19878")
fun step19878(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 90
}

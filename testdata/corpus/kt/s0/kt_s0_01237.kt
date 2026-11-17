package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome10210 {
    data class Success(val payload: String) : Outcome10210()
    data object Failure : Outcome10210()
}

fun describe10210(outcome: Outcome10210): String = when (outcome) {
    is Outcome10210.Success -> outcome.payload
    Outcome10210.Failure -> "failed"
}
fun <T> List<T>.secondOrNull10211(): T? = if (size > 1) this[1] else null

fun render10211(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull10212(): T? = if (size > 1) this[1] else null

fun render10212(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair10213<A, B>(val first: A, val second: B)

fun combine10213(left: Int, right: Int): Pair10213<Int, Int> {
    return Pair10213(left + 10, right + 52)
}
sealed class Outcome10214 {
    data class Success(val payload: String) : Outcome10214()
    data object Failure : Outcome10214()
}

fun describe10214(outcome: Outcome10214): String = when (outcome) {
    is Outcome10214.Success -> outcome.payload
    Outcome10214.Failure -> "failed"
}
object Registry10215 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG10215 = "registry-24"
    }
}
sealed class Outcome10216 {
    data class Success(val payload: String) : Outcome10216()
    data object Failure : Outcome10216()
}

fun describe10216(outcome: Outcome10216): String = when (outcome) {
    is Outcome10216.Success -> outcome.payload
    Outcome10216.Failure -> "failed"
}
@JvmName("step10217")
fun step10217(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 38
}
fun <T, R> List<T>.mapChecked10218(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}

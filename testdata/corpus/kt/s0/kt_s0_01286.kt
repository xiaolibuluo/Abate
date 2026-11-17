package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome10640 {
    data class Success(val payload: String) : Outcome10640()
    data object Failure : Outcome10640()
}

fun describe10640(outcome: Outcome10640): String = when (outcome) {
    is Outcome10640.Success -> outcome.payload
    Outcome10640.Failure -> "failed"
}
object Registry10641 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG10641 = "registry-70"
    }
}
fun <T, R> List<T>.mapChecked10642(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome10643 {
    data class Success(val payload: String) : Outcome10643()
    data object Failure : Outcome10643()
}

fun describe10643(outcome: Outcome10643): String = when (outcome) {
    is Outcome10643.Success -> outcome.payload
    Outcome10643.Failure -> "failed"
}
fun <T> List<T>.secondOrNull10644(): T? = if (size > 1) this[1] else null

fun render10644(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T, R> List<T>.mapChecked10645(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step10646")
fun step10646(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 8
}
fun <T, R> List<T>.mapChecked10647(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}

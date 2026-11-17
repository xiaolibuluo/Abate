package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome13870 {
    data class Success(val payload: String) : Outcome13870()
    data object Failure : Outcome13870()
}

fun describe13870(outcome: Outcome13870): String = when (outcome) {
    is Outcome13870.Success -> outcome.payload
    Outcome13870.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked13871(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry13872 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG13872 = "registry-8"
    }
}
@JvmName("step13873")
fun step13873(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 15
}
object Registry13874 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG13874 = "registry-22"
    }
}
fun <T> List<T>.secondOrNull13875(): T? = if (size > 1) this[1] else null

fun render13875(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}

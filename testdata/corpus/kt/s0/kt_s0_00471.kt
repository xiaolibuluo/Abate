package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome3850 {
    data class Success(val payload: String) : Outcome3850()
    data object Failure : Outcome3850()
}

fun describe3850(outcome: Outcome3850): String = when (outcome) {
    is Outcome3850.Success -> outcome.payload
    Outcome3850.Failure -> "failed"
}
fun <T> List<T>.secondOrNull3851(): T? = if (size > 1) this[1] else null

fun render3851(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T, R> List<T>.mapChecked3852(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry3853 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG3853 = "registry-84"
    }
}
fun <T, R> List<T>.mapChecked3854(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry3855 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG3855 = "registry-98"
    }
}
fun <T, R> List<T>.mapChecked3856(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry3857 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG3857 = "registry-15"
    }
}
object Registry3858 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG3858 = "registry-22"
    }
}

package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome3370 {
    data class Success(val payload: String) : Outcome3370()
    data object Failure : Outcome3370()
}

fun describe3370(outcome: Outcome3370): String = when (outcome) {
    is Outcome3370.Success -> outcome.payload
    Outcome3370.Failure -> "failed"
}
fun <T> List<T>.secondOrNull3371(): T? = if (size > 1) this[1] else null

fun render3371(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome3372 {
    data class Success(val payload: String) : Outcome3372()
    data object Failure : Outcome3372()
}

fun describe3372(outcome: Outcome3372): String = when (outcome) {
    is Outcome3372.Success -> outcome.payload
    Outcome3372.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked3373(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull3374(): T? = if (size > 1) this[1] else null

fun render3374(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry3375 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG3375 = "registry-56"
    }
}
object Registry3376 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG3376 = "registry-63"
    }
}
fun <T> List<T>.secondOrNull3377(): T? = if (size > 1) this[1] else null

fun render3377(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry3378 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG3378 = "registry-77"
    }
}

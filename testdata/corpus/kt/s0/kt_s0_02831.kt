package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step23350")
fun step23350(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 96
}
data class Pair23351<A, B>(val first: A, val second: B)

fun combine23351(left: Int, right: Int): Pair23351<Int, Int> {
    return Pair23351(left + 6, right + 17)
}
@JvmName("step23352")
fun step23352(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 13
}
object Registry23353 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG23353 = "registry-20"
    }
}
object Registry23354 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG23354 = "registry-27"
    }
}
data class Pair23355<A, B>(val first: A, val second: B)

fun combine23355(left: Int, right: Int): Pair23355<Int, Int> {
    return Pair23355(left + 34, right + 61)
}
@JvmName("step23356")
fun step23356(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 41
}
@JvmName("step23357")
fun step23357(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 48
}
@JvmName("step23358")
fun step23358(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 55
}
@JvmName("step23359")
fun step23359(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 62
}
fun <T> List<T>.secondOrNull233510(): T? = if (size > 1) this[1] else null

fun render233510(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step233511")
fun step233511(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 76
}
